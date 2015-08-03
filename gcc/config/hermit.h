/* Useful if you wish to make target-specific gcc changes. */
#undef TARGET_HERMIT
#define TARGET_HERMIT 1
 
/* Default arguments you want when running your
 *    i686-myos-gcc/x86_64-myos-gcc toolchain */
#define LIB_SPEC "-lc -lg -lm -lgloss" /* link against C standard libraries */
                                       /* modify this based on your needs */
 
/* Don't automatically add extern "C" { } around header files. */
#undef  NO_IMPLICIT_EXTERN_C
#define NO_IMPLICIT_EXTERN_C 1
 
/* Additional predefined macros. */
#undef TARGET_OS_CPP_BUILTINS
#define TARGET_OS_CPP_BUILTINS()      \
  do {                                \
    builtin_define ("__hermit__");      \
    builtin_define ("__unix__");      \
    builtin_assert ("system=hermit");   \
    builtin_assert ("system=unix");   \
    builtin_assert ("system=posix");   \
  } while(0);
