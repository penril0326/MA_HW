#!/usr/bin/perl
require "./validation/validate.pl";

# shorter than 8
sub testShorterThan8 {
    print "8-: 1: FAILED" if validate("1234567") != 0;
    print "8-: 2: FAILED" if validate("abc123") != 0;
    print "8-: 3: FAILED" if validate("aB12\$\#") != 0;
}

# 8 - 11
sub test8to11 {
    print "8-11: 1: FAILED" if validate("1233abCD!@#") != 1;
    print "8-11: 2: FAILED" if validate("abcd1234!@#") != 0;
    print "8-11: 3: FAILED" if validate("abcdEFGD12") != 0;
    print "8-11: 4: FAILED" if validate("aBcD1234++") != 1;
}

# 12 - 15
sub test12to15 {
    print "12-15: 1: FAILED" if validate("abcdeFGHIJ12345") != 1;
    print "12-15: 2: FAILED" if validate("abcdefghij12345") != 0;
    print "12-15: 3: FAILED" if validate("ABCDEFGHIJ12345") != 0;
    print "12-15: 4: FAILED" if validate("abcdefghij+-=*&") != 0;
    print "12-15: 5: FAILED" if validate("abcdEFGHIJ+-123") != 1;
}

# 16 - 19
sub test16to19 {
    print "16-19: 1: FAILED" if validate("ABCDEFGhijklmnopqrs") != 1;
    print "16-19: 2: FAILED" if validate("abcdefghiJKLMN12345") != 1;
    print "16-19: 3: FAILED" if validate("abcdefghiJKLMN123=-)") != 1;
    print "16-19: 4: FAILED" if validate("abcdefghijkl1234567") != 0;
}

# 20+
sub test20up {
    print "20+: 1: FAILED" if validate("abcdefghijklmnopqrstuvwxyz") != 1;
    print "20+: 2: FAILED" if validate("ABCDEFGHIJKLMNOPQRSTUVWXYZ") != 1;
    print "20+: 3: FAILED" if validate("12345678901234567890123") != 1;
    print "20+: 4: FAILED" if validate("!@#\$\%^&*()-+!@!#^&*()-+=") != 1;
}