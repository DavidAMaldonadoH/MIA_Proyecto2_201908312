grammar Gramatica;

options {
    language='Go';
}

@parser::header { 
    import "Proyecto2/Util"
    import "Proyecto2/Comandos" 
    import "strings"
}

start returns[ []util.Command list ]
    : commands EOF { $list = $commands.list }
    ;

commands returns[ []util.Command list ]
    : c1=comando { $list = append($list, $c1.c) } ( c2=comando { $list = append($list, $c2.c) } )*
    ;

comando returns[ util.Command c ]
    : EXEC parameters   { $c = comandos.NewExec($parameters.params) }
    | MKDISK parameters { $c = comandos.NewMkdisk($parameters.params) }
    | RMDISK parameters { $c = comandos.NewRmdisk($parameters.params) }
    | FDISK parameters  { $c = comandos.NewFdisk($parameters.params) }
    | MOUNT parameters  { $c = comandos.NewMount($parameters.params) }
    | MKFS parameters   { $c = comandos.NewMkfs($parameters.params) }
    | LOGIN parameters  { $c = comandos.NewLogin($parameters.params) }
    | LOGOUT            { $c = comandos.NewLogout() }
    | MKGRP parameters  { $c = comandos.NewMkgroup($parameters.params) }
    | RMGRP parameters  { $c = comandos.NewRmgroup($parameters.params) }
    | MKUSER parameters { $c = comandos.NewMkuser($parameters.params) }
    | RMUSR parameters  { $c = comandos.NewRmusr($parameters.params) }
    | MKFILE parameters { $c = comandos.NewMkfile($parameters.params) }
    | MKDIR parameters  { $c = comandos.NewMkdir($parameters.params) }
    | PAUSE             { $c = comandos.NewPause() }
    | REP parameters    { $c = comandos.NewRep($parameters.params) }
    ;

fit_type returns[ string fit ]
	: BF { $fit = "bf" }
	| FF { $fit = "ff" }
	| WF { $fit = "wf" }
	;

unit_type returns[ string u ]
    : KILO { $u = "k" }
    | MEGA { $u = "m" }
    | BYTE { $u = "b" }
    ;

part_type returns[ string t ]
    : PRIMARY { $t = "p" }
    | EXTENDED { $t = "e" }
    | LOGIC { $t = "l" }
	;

format_type returns [ string t ]
    : FAST { $t = "fast" }
    | FULL { $t = "full" }
    ;

texto returns[ string val ]
	: STRING { $val = strings.ReplaceAll($STRING.text, "\"", "") }
	| DIR { $val = $DIR.text }
    | id1=ID { $val = $id1.text }
    | id1=ID '.' id2=ID { $val = $id1.text + "." + $id2.text }
    | ENTERO {$val = $ENTERO.text }
	;

parameters returns[ []util.Parameter params ]
	: p1=param { $params = append($params, $p1.p) } ( p2=param { $params = append($params, $p2.p) } )* 
	;
	
param returns [ util.Parameter p ]
	: '-' SIZE '=' ENTERO { $p = util.NewParameter("size", $ENTERO.int, true) }
	| '-' FIT '=' fit_type { $p = util.NewParameter("fit", $fit_type.fit, false) }
	| '-' UNIT '=' unit_type { $p = util.NewParameter("unit", $unit_type.u, false) }
	| '-' PATH '=' texto { $p = util.NewParameter("path", $texto.val, true) }
    | '-' TYPE '=' part_type { $p = util.NewParameter("type", $part_type.t, false) }
    | '-' TYPE '=' format_type { $p = util.NewParameter("type", $format_type.t, false) }
    | '-' NAME '=' texto { $p = util.NewParameter("name", $texto.val, true) }
    | '-' IDENT '=' ID { $p = util.NewParameter("id", $ID.text, true) }
    | '-' USUARIO '=' texto { $p = util.NewParameter("usuario", $texto.val, true) }
    | '-' PASSWORD '=' texto { $p = util.NewParameter("password", $texto.val, true) }
    | '-' PWD '=' texto { $p = util.NewParameter("pwd", $texto.val, true) }
    | '-' GRP '=' texto { $p = util.NewParameter("grp", $texto.val, true) }
    | '-' RP { $p = util.NewParameter("r", nil, false) }
    | '-' CONT '=' texto { $p = util.NewParameter("cont", $texto.val, false) }
    | '-' RUTA '=' texto { $p = util.NewParameter("ruta", $texto.val, false) }
	;

// commands
EXEC: E X E C;
MKDISK: M K D I S K;
RMDISK: R M D I S K;
FDISK: F D I S K;
MOUNT: M O U N T;
MKFS: M K F S;
LOGIN: L O G I N;
LOGOUT: L O G O U T;
MKGRP: M K G R P;
RMGRP: R M G R P;
MKUSER: M K U S R;
RMUSR: R M U S R;
MKFILE: M K F I L E;
MKDIR: M K D I R;
PAUSE: P A U S E;
REP: R E P;

// parametros
SIZE: S I Z E;
FIT: F I T;
UNIT: U N I T;
PATH: P A T H;
TYPE: T Y P E;
NAME: N A M E;
IDENT: I D;
USUARIO: U S U A R I O;
PASSWORD: P A S S W O R D;
PWD: P W D;
GRP: G R P;
RP: R;
CONT: C O N T;
RUTA: R U T A;

// valores
BF:         B F;
FF:         F F;
WF:         W F;
KILO:       K;
MEGA:       M;
BYTE:       B;
PRIMARY:    P;
EXTENDED:   E;
LOGIC:      L;
FAST:       F A S T;
FULL:       F U L L;

// tokens
ENTERO: '-'?[0-9]+;
STRING: '"'~["]*'"';
ID: ([a-zA-Z]|[0-9]|'_')+;
DIR: ('/'~["|\n]+)+;
COMS: '#' .*? '\n' -> skip; // comentario una línea
COMD: '#*' .*? '*#' -> skip; // comentario multiple líneas
WS: [ \r\n\t]+ -> skip; // espacios

fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];