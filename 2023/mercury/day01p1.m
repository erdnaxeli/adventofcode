% vim: ft=mercury ff=unix ts=4 sw=4 et

:- module day01p1.
:- interface.
:- import_module io.
:- pred main(io::di, io::uo) is cc_multi.
:- implementation.

:- import_module char.
:- import_module int.
:- import_module list.
:- import_module string.

:- import_module file.

:- pred compute_solution(list(list(char))::in, int::out) is cc_multi.
compute_solution(Lines, Result) :-
    (
        compute_solution_2(Lines, 0, Result)
    ;
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
    append(NotDigits1, Rest1, Line),
    Rest1 = [ FirstDigitS | Rest2 ],
    append(_, Rest3, Rest2),
    Rest3 = [ LastDigitS | NotDigits2],
    not any_digit(NotDigits1),
    not any_digit(NotDigits2),
    to_int(from_char_list([FirstDigitS]), FirstDigit),
    to_int(from_char_list([LastDigitS]), LastDigit),
    Value = FirstDigit*10 + LastDigit.

:- pred any_digit(list(char)::in) is semidet.
any_digit([]) :- fail.
any_digit([ X | Xs ]) :-
    (
        to_int(from_char_list([X]), _)
    ;
        any_digit(Xs)
    ).

main(!IO) :-
    file.read_file("day01.txt", Lines, !IO),
    compute_solution(Lines, Solution),
    io.format("%d", [i(Solution)], !IO).
