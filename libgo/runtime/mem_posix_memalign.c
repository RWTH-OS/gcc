#include <errno.h>

#include "runtime.h"
#include "arch.h"
#include "malloc.h"

#define PAGE_FLOOR(addr)	(((addr) + PageSize - 1) & ~(PageSize-1))

void*
runtime_SysAlloc(uintptr n, uint64 *stat)
{
	USED(stat);
	void *p;

	mstats.sys += n;
#ifdef __hermit__
	p = malloc(n+PageSize);
	p = (void*) PAGE_FLOOR((size_t) p);
#else
	errno = posix_memalign(&p, PageSize, n);
	if (errno > 0) {
		perror("posix_memalign");
		exit(2);
	}
#endif
	return p;
}

void
runtime_SysUsed(void *v, uintptr n)
{
	USED(v);
	USED(n);
}

void
runtime_SysUnused(void *v, uintptr n)
{
	USED(v);
	USED(n);
	// TODO(rsc): call madvise MADV_DONTNEED
}

void
runtime_SysFree(void *v, uintptr n, uint64 *stat)
{
	USED(stat);
	mstats.sys -= n;
	free(v);
}

// add dummy pointer for "runtime/malloc.goc"
char* _end = NULL;

void*
runtime_SysReserve(void *v, uintptr n, bool* reserved)
{
	USED(v);
	USED(reserved);
	return runtime_SysAlloc(n, NULL);
}

void
runtime_SysMap(void *v, uintptr n, bool reserved, uint64 *stat)
{
	USED(v);
	USED(n);
	USED(stat);
	USED(reserved);
}

void
runtime_SysFault(void *v, uintptr n)
{
	USED(v);
	USED(n);
}
