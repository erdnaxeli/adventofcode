% vim: ft=mercury ff=unix ts=4 sw=4 et

:- module day01.
:- interface.
:- import_module io.
:- pred main(io::di, io::uo) is det.
:- implementation.

:- import_module char.
:- import_module int.
:- import_module list.
:- import_module string.

:- pred compute_solution(list(list(char))::in, int::out) is det.
compute_solution(Lines, Result) :-
    compute_solution_2(Lines, 0, Result).

:- pred compute_solution_2(list(list(char))::in, int::in, int::out) is det.
compute_solution_2([ Line | Lines ], Acc, Result) :-
    compute_value(Line, Value),
    compute_solution_2(Lines, Acc + Value, Result).

compute_solution_2([], Acc, Result) :-
    Result = Acc.

:- pred compute_value(list(char)::in, int::out) is det.
compute_value(Line, Value) :-
    compute_first_digit(Line, First_digit),
    compute_first_digit(reverse(Line), Second_digit),
    Value = First_digit*10 + Second_digit.

:- pred compute_first_digit(list(char)::in, int::out) is det.
compute_first_digit([], 0).
compute_first_digit([ X | Xs ], Digit) :-
    ( if
        to_int(char_to_string(X), Digit0)
    then
        Digit = Digit0
    else
        compute_first_digit(Xs, Digit)
    ).

:- pred read_file(io.text_input_stream::in, list(list(char))::out, io::di, io::uo).
read_file(Stream, Lines, !IO) :-
    io.read_line(Stream, Result, !IO),
    (
        Result = eof,
        Lines = []
    ;
        Result = error(_),
        Lines = []
    ;
        Result = ok(Line),
        read_file(Stream, Other_lines, !IO),
        Lines = cons(Line, Other_lines)
    ).

main(!IO) :-
    io.open_input("day01.txt", Result, !IO),
    (
        Result = ok(Stream),
        read_file(Stream, Lines, !IO),
        io.close_input(Stream, !IO),
        compute_solution(Lines, Solution),
        io.format("%d", [i(Solution)], !IO)
    ;
        Result = error(Error),
        Msg = io.error_message(Error),
        io.write_string(Msg, !IO)
    ).
