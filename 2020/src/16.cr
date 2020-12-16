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

  def self.find_rules_positions(tickets, rules)
    candidate_positions = Hash(String, Hash(Int32, Int32)).new { |h, k| h[k] = Hash(Int32, Int32).new { 0 } }

    # count how many tickets are valide for each rule position
    tickets.each do |ticket|
      rules.each do |rule|
        ticket.each_with_index do |value, idx|
          if rule.ranges.any? &.includes?(value)
            candidate_positions[rule.field][idx] += 1
          end
        end
      end
    end

    # While there are rules with unknown position:
    # * we search the next rule which has only one possible position
    # * we store this position for this rule
    final_positions = Hash(String, Int32).new
    while candidate_positions.size > 0
      found_position = nil
      found_field, _ = candidate_positions.find do |_, positions|
        positions.select { |p, _| !final_positions.has_value?(p) }.count do |candidate, count|
          if count == tickets.size
            found_position = candidate
          end
        end == 1
      end.not_nil!

      final_positions[found_field] = found_position.not_nil!
      candidate_positions.delete(found_field)
    end

    final_positions
  end

  INPUT_DAY15 = File.read("./inputs/16.txt")

  def self.day16p1
    rules_input, _, tickets_input = INPUT_DAY15.split("\n\n")
    rules = read_ticket_rules(rules_input)
    tickets = read_tickets(tickets_input.lines[1..])

    tickets.flat_map { |ticket| find_invalid_ticket_field(ticket, rules) }.sum
  end

  def self.day16p2
    rules_input, own_ticket_input, tickets_input = INPUT_DAY15.split("\n\n")
    rules = read_ticket_rules(rules_input)
    tickets = read_tickets(tickets_input.lines[1..]).select do |ticket|
      find_invalid_ticket_field(ticket, rules).size == 0
    end

    own_ticket = own_ticket_input.lines[1].split(',').map &.to_i

    result = 1_u64
    find_rules_positions(tickets, rules).each do |field, position|
      if field.starts_with?("departure")
        result *= own_ticket[position].to_u64
      end
    end

    result
  end
end
