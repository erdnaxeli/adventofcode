module NiceString
  extend self

  def nice?(s) : Bool
    couples = s.each_char.cons(2).map(&.join).to_a
    counter = s.each_char.tally
    vowels_count = ->{ "aeiou".each_char.sum { |c| counter[c]? || 0 } }
    double_letter = ->{ couples.any? { |c| c[0] == c[1] } }
    followings = ->{ couples.any? { |c| ["ab", "cd", "pq", "xy"].includes? c } }

    vowels_count.call >= 3 && double_letter.call && !followings.call
  end

  def nice2?(s) : Bool
    if s.size < 4
      return false
    end

    pair_repeated = s.matches? /(..).*\1/
    letter_repeated = s.matches? /(.).\1/
    pair_repeated && letter_repeated
  end
end
