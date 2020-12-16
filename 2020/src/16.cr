module Aoc2020
  record(
    TicketRule,
    field : String,
    ranges : Array(Range(Int32, Int32)),
  )

  alias Ticket = Array(Int32)

  def self.read_ticket_rules(input)
    input.each_line.compact_map do |l|
      if result = /(?<field>[a-z ]+): (?<r11>\d+)-(?<r12>\d+) or (?<r21>\d+)-(?<r22>\d+)/.match(l)
        TicketRule.new(
          field: result["field"],
          ranges: [
            (result["r11"].to_i..result["r12"].to_i),
            (result["r21"].to_i..result["r22"].to_i),
          ]
        )
      end
    end.to_a
  end

  def self.read_tickets(input) : Array(Ticket)
    input.map do |l|
      l.split(',').map &.to_i
    end
  end

  def self.find_invalid_ticket_field(ticket, rules)
    invalid_fields = Array(Int32).new

    ticket.each do |value|
      if rules.each.flat_map(&.ranges).any? &.includes?(value)
        next
      else
        invalid_fields << value
      end
    end

    invalid_fields
  end

  INPUT_DAY15 = File.read("./inputs/16.txt")

  def self.day16p1
    rules_input, _, tickets_input = INPUT_DAY15.split("\n\n")
    rules = read_ticket_rules(rules_input)
    tickets = read_tickets(tickets_input.lines[1..])

    tickets.flat_map { |ticket| find_invalid_ticket_field(ticket, rules) }.sum
  end
end

puts Aoc2020.day16p1
