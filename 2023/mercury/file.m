% vim: ft=mercury ff=unix ts=4 sw=4 et

:- module file.
:- interface.
:- import_module char.
:- import_module list.
:- import_module io.
:- import_module string.
:- pred read_file(string::in, list(list(char))::out, io::di, io::uo) is det.
:- implementation.

read_file(Filename, Lines, !IO) :-
    io.open_input(Filename, Result, !IO),
    (
        Result = ok(Stream),
        read_lines(Stream, Lines, !IO)
    ;
        Result = error(Error),
        Msg = io.error_message(Error),
        io.write_string(Msg, !IO),
        Lines = []
    ).

:- pred read_lines(io.text_input_stream::in, list(list(char))::out, io::di, io::uo) is det.
read_lines(Stream, Lines, !IO) :-
    io.read_line(Stream, Result, !IO),
    (
        Result = eof,
        Lines = []
    ;
        Result = error(_),
        Lines = []
    ;
        Result = ok(Line),
        read_lines(Stream, Lines_2, !IO),
        Lines = cons(Line, Lines_2)
    ).
