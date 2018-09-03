uses crt,myio,windows;

{$I psedit.inc}

const
 ver = '.50/alpha';

 n_attr        = $17; {normal text attribute}
 status_attr   = $3f; {status line attribute}
 field_attr    = $71; {color for input fields}
 title_attr    = $71; {title line color}
 lightbar_attr = $7F; {lightbar color}
 i_attr        = $1f; {important stuff}

 c_u = $0d0e;
 c_b = $000f;
 c_n = $2000;

procedure nl(i:word);
var
 j:integer;
begin
for j := 1 to i do writeln;
end;

procedure setcur(c:word);forward;

procedure showlist(choices:choicetype;starty:integer);
var
i:integer;
 begin
i:= 1;
textattr:= n_attr;
gotoxy(1,starty+1);
 while(choices[i][1]<>#0) do
  begin                                 {shows list}
   writeln('   ',choices[i],'  ');
   inc(i);
  end;
end;

function lightbar(choices:choicetype) : integer;
var
 tempattr : byte;
 hold:char;
 i : integer;
 barpos :integer;
 done:boolean;
 starty : integer;
 prevbar:integer;

begin
barpos := 1;
done := false;
tempattr := textattr;

setcur(c_n);

starty := wherey;
showlist(choices,starty);

i:= 0;
repeat
inc(i);
until choices[i] = '';


 repeat
gotoxy(1,starty);
showlist(choices,starty);

gotoxy(1,starty+barpos);
write(' ');
textattr := lightbar_attr;
writeln('  ',choices[barpos],'  ');

hold := readkey;

 if hold = #0 then
 begin
  hold := readkey;
   case hold of

   #72: if barpos > 0+1 then
         begin
          prevbar := barpos;
          dec(barpos);
         end;
   #80: if barpos < i-1 then
         begin
          prevbar := barpos;
          inc(barpos);
        end;
   #68:begin {done := true} end;
  end;                  {case}
 end;                   {if hold=0}

 if hold = #13 then
  begin
   lightbar := barpos;
   done := true;
  end;

{lightbar}


 until done;
textattr := tempattr;
setcur(c_u);
gotoxy(1,starty+i);
end;

procedure setcur(c:word);assembler;
asm
 mov ah,1
 mov cx,c
 int 10h
end;


procedure update_status(p:string);
var
 smin, smax : word;
 sx, sy, st : byte;
begin
setcur(c_n);

smin := windmin;
smax := windmax;

sx := wherex;
sy := wherey;
st := textattr;

window(1,25,80,25);
textattr := status_attr;
clreol;
c_write(p);

if(p[length(p)]=#255)then
 begin
  if readkey=#0 then readkey;
  setcur(c_u);
 end;

windmin := smin;
windmax := smax;

gotoxy(sx,sy);

textattr := st;
setcur(c_u);
end;


procedure listitems(it:itemlist;numitems:byte);
var
 i:integer;
 begin
 if numitems=0 then exit;
 for i := 1 to numitems do begin write('  ');write_item(it[i]);writeln;end;
 end;

procedure clearchoices(var c:choicetype);
var
 i:integer;
begin
for i := 1 to maxchoices+1 do c[i] := ''#0;
end;

procedure setcnames(var c:choicetype);
begin
c[1] := 'Alis';
c[2] := 'Myau';
c[3] := 'Odin';
c[4] := 'Noah';
c[5] := ''#0;
end;

procedure setcmenu(var c:choicetype);
begin
c[1] := 'Player Editor';
c[2] := ' Item Editor ';
c[3] := 'Change Meseta';
c[4] := ''#0;
end;


procedure writestats(pr:playerrec);
begin
with pr do
 begin
 writeln;
 writeln('            HP =',curhp:6,'/',maxhp:5,'  |  exp    =' ,exp:6,  '  |  level   =',level:6);
 writeln('            MP =',curmp:6,'/',maxmp:5,'  |  attack =',attack:6,'  |  defense =',defense:6);
 writeln;

 writeln('                                     state = ',state:2);
 writeln('                             combat spells = ',cspells:2);
 writeln('                         non-combat spells = ',ncspells:2);
 writeln;
 write('    weapon = ');write_item(weapon);writeln;
 write('    armor  = ');write_item(armor); writeln;
 write('    shield = ');write_item(shield);writeln;

 end;
end;

procedure write_snames(h:headertype;var c:choicetype);

 procedure write_stupidname(num:word;var c:choicetype);
 var
  i,k : integer;
  ch : char;

 begin
 textattr := n_attr;
 write('[');
 textattr := i_attr;
 for k := 1 to 5 do begin
   ch := charset[ord(h[nameofs[num]])-138-65];
    {chr(ord(h[ofs])-138); = raw ascii}
   if(chr(ord(h[nameofs[num]])-138)<>'6')then
    begin
    write(ch);
    c[num] := c[num]+ch;
    end;
   inc(nameofs[num],2); {trick pascal into changing our constant.  ha!}
  end;
  dec(nameofs[num],10); {put the constant back, or else!}
 textattr := n_attr;
 write(']');
 textattr := $18;
 if ord(h[delofs[num]])=0 then write(' [deleted]');
 writeln;
 end;

var
 a:byte;
 i:integer;

begin
{savegame title offsets =
 $2a = game #1
 $4e = game #2
 $72 = game #3
 $96 = game #4
 $ba = game #5

 $101-$105 = savegame exists

 all offsets are + 100h in actual savegame
}

a := textattr;
write(' ');
textattr := status_attr;
writeln('  saved games:  ');
clearchoices(c);
for i := 1 to 5 do c[i] := '';

textattr := n_attr;write('  [1]  : ');write_stupidname(1,c);
textattr := n_attr;write('  [2]  : ');write_stupidname(2,c);
textattr := n_attr;write('  [3]  : ');write_stupidname(3,c);
textattr := n_attr;write('  [4]  : ');write_stupidname(4,c);
textattr := n_attr;write('  [5]  : ');write_stupidname(5,c);
textattr := a;
end;

procedure main;
var
 fn : string;
 temps : string;
 gamenum : byte;
 g_ofs   : word;

 header : headertype;

 meseta : word;
 numitems : byte;
 inventory : itemlist;

 f : file;
 pr : playerrec;

 choices,savenames : choicetype;
 choice,ch2 : integer;

 sx, sy : byte;

 cstr : string;

begin
fn := 'phanstar.sav';
temps := '';
textattr := title_attr;
clreol;
write(' PSEDIT version ',ver,' by Mike Pontillo and Kevin Nishi');
update_status(' [F10] = Quit ');
window(1,2,80,24);
textattr := n_attr;
clrscr;
writeln;

writeln('  enter the name of savegame file');
write(':');
textattr := field_attr;
myreadln(fn,78);
if exitflag then exit;
textattr := n_attr;

if not fileexists(fn) then
 begin
 textattr := $9F;
 nl(2);
 writeln('FATAL ERROR:  savegame file not found.');
 update_status(s_osret);
 exit;
 end;

assign(f,fn);
reset(f,1);

nl(2);
seek(f,$100);
blockread(f,header,sizeof(header));

repeat
clrscr;
writeln;
write_snames(header,savenames);
{

FUCK!  i'll get this working later.

lightbar(savenames);
}
writeln;
write(' which savegame number (1-5), ([blank] = 1, or [F10] to exit)? ');
validchars := '12345';
textattr := field_attr;
myreadln(temps,1);
if temps = '' then
 begin
  temps := '1';
  write('1');
 end;
if exitflag then exit;
gamenum := strtoint(temps);
validchars := validdefs;
textattr := n_attr;

gamenum := gamenum - 1;
g_ofs := $500 + ($400*gamenum);
nl(2);
sx := wherex;
sy := wherey;

nl(6);
update_status(' editing savegame #'+inttostr(gamenum+1));

gotoxy(wherex,wherey-4);
textattr := $19;
box(2,31,wherey,51,wherey+6);
textattr := n_attr;
setcmenu(choices);
ch2 := lightbar(choices);
window(1,2,80,24);
gotoxy(sx,sy);


seek(f,g_ofs+(sizeof(pr)*4)+128);
blockread(f,inventory,sizeof(inventory));

seek(f,g_ofs+(sizeof(pr)*4)+160);
blockread(f,meseta,2);
blockread(f,numitems,1);


case ch2 of
1:begin

  clrscr;
  writeln;

  textattr := $1f;
  writeln(' which character do you want to edit?');
  textattr := n_attr;

  nl(9);

  sx := wherex;
  sy := wherey;

  gotoxy(wherex,wherey-7);

  textattr := $19;
  box(2,35,wherey,46,wherey+7);
  textattr := n_attr;
  setcnames(choices);
  choice := lightbar(choices);
  cstr := '['+choices[choice]+']';
  window(1,2,80,24);
  gotoxy(sx,sy);


  dec(choice);

  {
  item #1 = g_ofs+(sizeof(pr)*4)+128

  meseta = g_ofs+(sizeof(pr)*4)+160
  numitems = g_ofs+(sizeof(pr)*4)+162
  }


  clrscr;
  writeln;

  seek(f,g_ofs+(sizeof(pr)*choice));


  blockread(f,pr,sizeof(pr));

  nl(2);
  center(0,cstr);
  writestats(pr);

  update_status(s_anykey);
 end;

2:begin
  clrscr;
  writeln(' you currently have the following items in your inventory:');
  listitems(inventory,numitems);
  writeln;
  writeln(' numitems = ', numitems);
  update_status(s_anykey);
 end;
3:begin
  clrscr;
  writeln(' meseta = ',meseta);
  update_status(s_anykey);
 end;
end;
update_status(' [F10]=exit');

until true=false;

end;

begin
 asm
  mov ax, 3h
  int 10h
 end;
main;
update_status(s_osret);
 asm
  mov ax, 3h
  int 10h
 end;
end.