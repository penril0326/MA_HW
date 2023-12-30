#!/usr/bin/perl

# Passwords must be at least 8 characters long.
# Between 8-11: requires mixed case letters, numbers and symbols
# Between 12-15: requires mixed case letters and numbers
# Between 16-19: requires mixed case letters
# 20+: any characters desired

sub validate {
    $passwd = shift;
    if (length($passwd) < 8) {
        return 0;
    } elsif ($passwd =~ /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!-.:-@])[a-zA-Z!-.:-@\d]{8,11}$/) {
        return 1;
    } elsif ($passwd =~ /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z!-.:-@\d]{12,15}$/) {
        return 1;
    } elsif ($passwd =~ /^(?=.*[a-z])(?=.*[A-Z])[a-zA-Z!-.:-@\d]{16,19}$/) {
        return 1;
    } elsif (length($passwd) >= 20) {
        return 1;
    } else {
        return 0;
    }
}

1;