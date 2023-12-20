#include <pthread.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <ctype.h>
#include "bmlib_runtime.h"
#include "bmlib_internal.h"
#include "message.h"
#include "gflags/gflags.h"

DEFINE_int32(stack_size, 16, "stack size , KBytes, minimal 16KBytes");
DEFINE_int32(thread_num, 1, "how many threads to be created");
DEFINE_int32(thread_time, 1, "how many seconds thread to sleep");
DECLARE_bool(help);
DECLARE_bool(helpshort);

#define handle_error_en(en, msg) \
	do { errno = en; perror(msg); exit(EXIT_FAILURE); } while (0)

#define handle_error(msg) \
	do { perror(msg); exit(EXIT_FAILURE); } while (0)

struct thread_info {    /* Used as argument to thread_start() */
	pthread_t thread_id;        /* ID returned by pthread_create() */
	int       thread_num;       /* Application-defined thread # */
	char     *argv_string;      /* From command-line argument */
};

/* Thread start function: display address near top of our stack,
	 and return upper-cased copy of argv_string */

	static void *
thread_start(void *arg)
{
	struct thread_info *tinfo = (struct thread_info*)arg;
	char *uargv, *p;

	//printf("Thread %d: top of stack near %p; argv_string=%s\n",
	//		tinfo->thread_num, &p, tinfo->argv_string);

	uargv = strdup(tinfo->argv_string);
	if (uargv == NULL)
		handle_error("strdup");

	for (p = uargv; *p != '\0'; p++)
		*p = toupper(*p);
	sleep(FLAGS_thread_time);
	return uargv;
}

	int
main(int argc, char *argv[])
{
	int s, tnum, thread_num = 0;
	struct thread_info *tinfo = NULL;
	pthread_attr_t attr;
	int stack_size = 0;
	void *res;

	/* get and validate flags*/
	gflags::SetUsageMessage("command line brew\n"
			"usage: bm_pthread_create [--thread_num=1] [--stack_size=16]"
			"thread_num:\n"
			"  how many threads to be created.\n"
			"stack_size:\n"
			"  stack size of one thread, KBytes.");

	gflags::ParseCommandLineNonHelpFlags(&argc, &argv, true);
	if (FLAGS_help) {
		FLAGS_help = false;
		FLAGS_helpshort = true;
	}
	gflags::HandleCommandLineHelpFlags();

	stack_size = FLAGS_stack_size * 1024;
	thread_num = FLAGS_thread_num;
	printf("FLAGS_thread_num = %d, FLAGS_stack_size=%d KB, FLAGS_thread_time=%d \
 second\n", FLAGS_thread_num, FLAGS_stack_size, FLAGS_thread_time);

	/* Initialize thread creation attributes */

	s = pthread_attr_init(&attr);
	if (s != 0)
		handle_error_en(s, "pthread_attr_init");

	s = pthread_attr_setstacksize(&attr, stack_size);
	if (s != 0)
		handle_error_en(s, "pthread_attr_setstacksize");

	/* Allocate memory for pthread_create() arguments */

	tinfo = (struct thread_info*)calloc(thread_num, sizeof(struct thread_info));
	if (tinfo == NULL)
		handle_error("calloc");

	/* Create one thread for each command-line argument */

	for (tnum = 0; tnum < thread_num; tnum++) {
		tinfo[tnum].thread_num = tnum + 1;
		tinfo[tnum].argv_string = (char *)"aaaaa";//argv[optind + tnum];

		/* The pthread_create() call stores the thread ID into
			 corresponding element of tinfo[] */

		s = pthread_create(&tinfo[tnum].thread_id, &attr,
				&thread_start, &tinfo[tnum]);
		if (s != 0) {
			handle_error_en(s, "pthread_create");
		}
	}

	/* Destroy the thread attributes object, since it is no
		 longer needed */

	s = pthread_attr_destroy(&attr);
	if (s != 0)
		handle_error_en(s, "pthread_attr_destroy");

	/* Now join with each thread, and display its returned value */

	for (tnum = 0; tnum < thread_num; tnum++) {
		s = pthread_join(tinfo[tnum].thread_id, &res);
		if (s != 0)
			handle_error_en(s, "pthread_join");

		//printf("Joined with thread %d; returned value was %s\n",
		//		tinfo[tnum].thread_num, (char *) res);
		free(res);      /* Free memory allocated by thread */
	}

	free(tinfo);
	exit(EXIT_SUCCESS);
}
