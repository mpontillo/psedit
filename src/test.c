/* Output from p2c 1.21alpha-07.Dec.93, the Pascal-to-C translator */
/* From input file "test.pas" */


#include <p2c/p2c.h>


main(argc, argv)
int argc;
Char *argv[];
{
  long i;

  PASCAL_MAIN(argc, argv);
  for (i = 1; i <= 10; i++) {
    println(i);
/* p2c: test.pas, line 5: Warning: Symbol 'PRINTLN' is not defined [221] */
  }

  exit(EXIT_SUCCESS);
}



/* End. */
