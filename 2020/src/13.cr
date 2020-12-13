require "big"

module Aoc2020
  def self.read_bus_planning(input)
    input.split(',').reject { |x| x == "x" }.map &.to_i
  end

  record BusGap, id : BigInt, gap : Int32

  def self.read_bus_planning_with_gap(input)
    input.split(',').each.with_index.reject { |x| x[0] == "x" }.map { |x| {x[0].to_big_i, x[1]} }.map do |x|
      BusGap.new(
        id: x[0],
        gap: x[1],
      )
    end.to_a
  end

  def self.find_best_bus(start, planning)
    min_start = nil
    bus_id = nil

    planning.each do |id|
      bus_start = ((start // id) + 1) * id

      if min_start.nil? || bus_start < min_start
        min_start = bus_start
        bus_id = id
      end
    end

    {bus_id, min_start}
  end

  def self.find_t_for_gap(planning)
    bus1 = planning.shift
    bus2 = planning.shift

    t = 0
    period = bus1.id

    loop do
      t += period

      if t % bus1.id == 0 && (t + bus2.gap) % bus2.id == 0
        if planning.size > 0
          period = period * bus2.id
          bus2 = planning.shift
        else
          return t
        end
      end
    end
  end

  INPUT_DAY13 = File.read("./inputs/13.txt").lines

  def self.day13p1
    start = INPUT_DAY13[0].to_i
    planning = read_bus_planning(INPUT_DAY13[1])
    bus_id, bus_start = find_best_bus(start, planning)

    if !bus_id.nil? && !bus_start.nil?
      bus_id * (bus_start - start)
    end
  end

  def self.day13p2
    planning = read_bus_planning_with_gap(INPUT_DAY13[1])
    find_t_for_gap(planning)
  end
end

pp Aoc2020.day13p2
