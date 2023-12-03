% vim: ft=mercury ff=unix ts=4 sw=4 et

:- module day01p1.
:- interface.
:- import_module io.
:- pred main(io::di, io::uo) is cc_multi.
:- implementation.

:- import_module bool.
:- import_module char.
:- import_module int.
:- import_module list.
:- import_module string.

:- import_module file.

:- pred compute_solution(list(list(char))::in, int::out) is multi.
compute_solution(Lines, Result) :-
    ( if
        compute_solution_2(Lines, 0, Result0)
    then
        Result = Result0
    else
        Result = -1
    ).

:- pred compute_solution_2(list(list(char))::in, int::in, int::out) is nondet.
compute_solution_2([ Line | Lines ], Acc, Result) :-
    compute_value(Line, Value),
    compute_solution_2(Lines, Acc + Value, Result).

compute_solution_2([], Acc, Result) :-
    Result = Acc.

:- pred compute_value(list(char)::in, int::out) is nondet.
compute_value(Line, Value) :-
    % a line is a list of not digits
    append(NotDigits1, Rest1, Line),

    % followed by a digit
    Rest1 = [ FirstDigitS | _ ],
    to_int(char_to_string(FirstDigitS), FirstDigit),
    not any_digit(NotDigits1),

    % then a digit (it can be the same, that's why we use Rest1) followed by not digits
    append(_, Rest2, Rest1),
    Rest2 = [ LastDigitS | NotDigits2],
    to_int(char_to_string(LastDigitS), LastDigit),
    not any_digit(NotDigits2),

    Value = FirstDigit*10 + LastDigit.

:- pred any_digit(list(char)::in) is semidet.
any_digit([]) :- fail.

any_digit([ X | Xs ]) :-
    (
        to_int(char_to_string(X), _)
    ;
        any_digit(Xs)
    ).

main(!IO) :-
    file.read_file("day01.txt", Lines, !IO),
    compute_solution(Lines, Solution),
    io.format("%d", [i(Solution)], !IO).
