const
 maxchoices = 10;
 charset : array[0..34] of char = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ,;.!?-"  ';

 nameofs : array[1..5] of word = ($2a,$4e,$72,$96,$ba);
 delofs : array[1..5] of word = ($101,$102,$103,$104,$105);

 s_osret = '                     - press any key to return to the os -'#255;
 s_anykey = ' * press any key to continue'#255;

type
 playerrec=record

  alive    : boolean;
  curhp,
  curmp    : byte;
  exp      : word;
  level,
  maxhp,
  maxmp,
  attack,
  defense,
  weapon,
  armor,
  shield,
  state,
  cspells,
  ncspells : byte;
 end;

 itemlist=array[1..32] of byte;

 headertype = array[0..$200] of char;

 choicetype = array[1..maxchoices+1] of string[40];

function itemnum(i:byte):string;
begin
 case i of
   0:itemnum := '[nothing]';
   1:itemnum := 'Wooden Cane';
   2:itemnum := 'Short Sword';
   3:itemnum := 'Iron Sword';
   4:itemnum := 'Wand';
   5:itemnum := 'Iron Fang';
   6:itemnum := 'Iron Axe';
   7:itemnum := 'Titanium Sword';
   8:itemnum := 'Ceramic Sword';
   9:itemnum := 'Needle Gun';
  10:itemnum := 'Silver Fang';
  11:itemnum := 'Heat Gun';
  12:itemnum := 'Light Saber';
  13:itemnum := 'Laser Gun';
  14:itemnum := 'Laconium Sword';
  15:itemnum := 'Laconium Axe';
  16:itemnum := 'Lithium Armor';
  17:itemnum := 'White Mantle';
  18:itemnum := 'Light Suit';
  19:itemnum := 'Iron Armor';
  20:itemnum := 'Thick Fur';
  21:itemnum := 'Zirconium Armor';
  22:itemnum := 'Diamond Armor';
  23:itemnum := 'Laconium Armor';
  24:itemnum := 'Fraid Mantle';
  25:itemnum := 'Lithium Shield';
  26:itemnum := 'Bronze Shield';
  27:itemnum := 'Iron Shield';
  28:itemnum := 'Ceramic Shield';
  29:itemnum := 'Glove';
  30:itemnum := 'Laser Shield';
  31:itemnum := 'Mirror Shield';
  32:itemnum := 'Lacomium Shield';
  33:itemnum := 'Land Rover';
  34:itemnum := 'Hovercraft';
  35:itemnum := 'Ice Digger';
  36:itemnum := 'Cola';
  37:itemnum := 'Burger';
  38:itemnum := 'Flute';
  39:itemnum := 'Flash';
  40:itemnum := 'Escaper';
  41:itemnum := 'Transfer';
  42:itemnum := 'Magic Hat';
  43:itemnum := 'Alsulin';
  44:itemnum := 'Polymtrl';
  45:itemnum := 'Dungeon Key';
  46:itemnum := 'Sphere';
  47:itemnum := 'Torch';
  48:itemnum := 'Prism';
  49:itemnum := 'Nuts';
  50:itemnum := 'Hapsby';
  51:itemnum := 'Road Pass';
  52:itemnum := 'Passport';
  53:itemnum := 'Compass';
  54:itemnum := 'Cake';
  55:itemnum := 'Letter';
  56:itemnum := 'Laconium Pot';
  57:itemnum := 'Magic Lamp';
  58:itemnum := 'Amber Eye';
  59:itemnum := 'Gas Shield';
  60:itemnum := 'Crystal';

            { not usable }
  61:itemnum := 'M System';
  62:itemnum := 'Miricle Key';
  63:itemnum := 'Zillion';
  64:itemnum := 'Secrets';
 end;
end;

procedure write_item(i:byte);
begin
write(itemnum(i));
end;
