module Aoc2020
  INPUT_DAY22 = File.read("./inputs/22.txt")

  def self.play(deck1, deck2)
    while deck1.size > 0 && deck2.size > 0
      card1 = deck1.shift
      card2 = deck2.shift

      puts "#{deck1.size} #{card1} <=> #{card2} #{deck2.size}"

      if card1 > card2
        deck1 << card1 << card2
      elsif card1 < card2
        deck2 << card2 << card1
      else
        raise "draw"
      end
    end

    if deck1.size > 0
      score(deck1)
    else
      score(deck2)
    end
  end

  def self.play2(deck1, deck2, i = 0)
    cache = Set(UInt64).new

    # puts "new game #{cache.size} #{i}"
    while deck1.size > 0 && deck2.size > 0
      if !cache.add?({deck1, deck2}.hash)
        # puts "nop"
        return -1, deck1
      end

      card1 = deck1.shift
      card2 = deck2.shift
      # puts "#{deck1.size} #{card1} <=> #{card2} #{deck2.size}"
      # sleep 100.milliseconds

      if card1 <= deck1.size && card2 <= deck2.size
        result, _ = play2(deck1[...card1].dup, deck2[...card2].dup, i+1)

        if result > 0
          deck2 << card2 << card1
        else
          deck1 << card1 << card2
        end
      elsif card1 > card2
        deck1 << card1 << card2
      elsif card1 < card2
        deck2 << card2 << card1
      else
        raise "draw"
      end
    end

    if deck1.size == 0
      return {1, deck2}
    else
      return {-1, deck1}
    end
  end

  def self.score(deck)
    deck.reverse.each_with_index(1).sum { |x| x[0] * x[1] }
  end

  def self.day22p1
    deck1, _, deck2 = INPUT_DAY22.partition("\n\n")
    play(deck1.lines[1..].map(&.to_i), deck2.lines[1..].map(&.to_i))
  end

  def self.day22p2
    deck1, _, deck2 = INPUT_DAY22.partition("\n\n")
    _, wining_deck = play2(deck1.lines[1..].map(&.to_i), deck2.lines[1..].map(&.to_i))
    score(wining_deck)
  end
end

puts Aoc2020.day22p2
