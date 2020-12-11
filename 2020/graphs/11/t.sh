convert \( 1/map.gif -coalesce -set page %[fx:w*2]x%[h]+0+0 -coalesce \) null: \( 2/map.gif -coalesce \) -gravity east -layers composite -set delay 15 -loop 0 result.gif
