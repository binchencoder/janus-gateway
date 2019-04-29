#!/usr/bin/perl
use strict;
use warnings;


print "Start generating localization resources\n";

sub readProto {
    my ($fileName) = @_;
    open (my $fin, $fileName) or die "failed to open file '$fileName' err $!";

    my @result;
    my $i = 0;
    while (my $row = <$fin>) {
        $result[$i] = $row;
        $i++;
    }
    close $fin;
    return @result;
}

sub parse {
    my (@lines) = @_;
    my $index = 0;
    my @result;
    my $locstr = "";

    foreach(@lines) {
        my $line = $_;
        if ($line =~ /^\s*(\S+)\s*=\s*(\d+).*;.*$/) {
            if (length($locstr) > 0) {
				$result[$index][0] = $2;
				$result[$index][1] = trim($locstr);
				$locstr = "";
				$index++;
			}
        } elsif ($line =~ /^\s*\/\/\@Trans(.+)$/) {
			$locstr .= $1;
		}
    }

    return @result;
}

sub parseVariableNameValue {
    my (@lines) = @_;
    my $index = 0;
    my @result;

    foreach(@lines) {
        my $line = $_;
        if ($line =~ /^\s*(\S+)\s*=\s*(\d+).*;.*#(.+)$/) {
            $result[$index][0] = $1;
            $result[$index][1] = $2;
            $index++;
        }
    }

    return @result;
}

sub generateLocVariables {
    my ($fileName, @rows) = @_;
    open (my $fout, '>', $fileName) or die "failed to open file '$fileName' err $!";
    foreach my $i (0 .. $#rows - 1) {
        print $fout "static $rows[$i][0]=$rows[$i][1];\n";
    }

    close $fout;
}


sub generateLocFile{
    my ($fileName, @rows) = @_;
    open (my $fout, '>', $fileName) or die "failed to open file '$fileName' err $!";

    print $fout "{\n";
    foreach my $i (0 .. $#rows) {
        # print "$rows[$i][0] = $rows[$i][1] \n";
        if ($i > 0) {
            print $fout ",\n";
        }

        print $fout "\"$rows[$i][0]\":\"$rows[$i][1]\"";
    }
    print $fout "\n}";

    close $fout;
}

sub trim {
	my $s = shift;
	$s =~ s/^\s+|\s+$//g;
	return $s
};

my @h = readProto('error.proto');
my @h1 = parse(@h);
generateLocFile("en.loc", @h1);

#@h1 = parseVariableNameValue(@h);
#generateLocVariables("var.js", @h1);

print "Done\n";





