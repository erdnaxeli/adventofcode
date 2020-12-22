use strict;
use warnings;
use v5.10;

sub parse_rules {
    my $input = shift @_;
    my %rules = {};

    foreach my $line (split(/\n/, $input)) {
        my ($n, $rule) = split(/: /, $line);
        my @branches = map { split(/ /, $_) } split(/ | ), $rule);
        $rules{$n} = \@branches;
    }

    return \%rules;
}

my $input = do { local(@ARGV, $/) = "./inputs/19.txt"; <> };
my ($rules_input, $messages) = split(/\n\n/, $input);

my %rules = parse_rules($rules_input)
say %rules;
