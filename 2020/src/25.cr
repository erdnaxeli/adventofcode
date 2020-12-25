module Aoc2020
  # Computes a one loop key.
  def self.compute_key(subject, value = 1)
    value *= subject
    value % 20201227
  end

  def self.find_loop_size(key)
    loop_size = 1
    candidate_key = compute_key(7)
    while candidate_key != key
      candidate_key = compute_key(7, candidate_key)
      loop_size += 1
    end

    loop_size
  end

  KEY1 = 2084668_u64
  KEY2 = 3704642_u64

  def self.day25p1
    encryption_key = 1_u64
    find_loop_size(KEY1).times { encryption_key = compute_key(KEY2, encryption_key) }
    encryption_key
  end
end
