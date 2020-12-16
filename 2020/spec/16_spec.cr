require "./spec_helper"

require "../src/16"

describe Aoc2020 do
  describe ".read_ticket_rules" do
    it "works" do
      input = %(class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50)
      rules = Aoc2020.read_ticket_rules(input)

      rules.size.should eq(3)
      rules[0].field.should eq("class")
      rules[1].field.should eq("row")
      rules[2].field.should eq("seat")
      rules[0].ranges.should eq([(1..3), (5..7)])
      rules[1].ranges.should eq([(6..11), (33..44)])
      rules[2].ranges.should eq([(13..40), (45..50)])
    end
  end

  describe ".read_tickets" do
    it "works" do
      input = %(7,3,47
40,4,50
55,2,20
38,6,12)
      Aoc2020.read_tickets(input.lines).should eq([
        [7, 3, 47],
        [40, 4, 50],
        [55, 2, 20],
        [38, 6, 12],
      ])
    end
  end

  describe ".find_invalid_ticket_field" do
    it "works 1" do
      ticket = [7, 3, 47]
      rules = [
        Aoc2020::TicketRule.new(field: "class", ranges: [(1..3), (5..7)]),
        Aoc2020::TicketRule.new(field: "row", ranges: [(6..11), (33..44)]),
        Aoc2020::TicketRule.new(field: "seat", ranges: [(13..40), (45..50)]),
      ]

      Aoc2020.find_invalid_ticket_field(ticket, rules).size.should eq(0)
    end

    it "works 2" do
      ticket = [40, 4, 50]
      rules = [
        Aoc2020::TicketRule.new(field: "class", ranges: [(1..3), (5..7)]),
        Aoc2020::TicketRule.new(field: "row", ranges: [(6..11), (33..44)]),
        Aoc2020::TicketRule.new(field: "seat", ranges: [(13..40), (45..50)]),
      ]

      Aoc2020.find_invalid_ticket_field(ticket, rules).should eq([4])
    end

    it "works 3" do
      ticket = [55, 2, 20]
      rules = [
        Aoc2020::TicketRule.new(field: "class", ranges: [(1..3), (5..7)]),
        Aoc2020::TicketRule.new(field: "row", ranges: [(6..11), (33..44)]),
        Aoc2020::TicketRule.new(field: "seat", ranges: [(13..40), (45..50)]),
      ]

      Aoc2020.find_invalid_ticket_field(ticket, rules).should eq([55])
    end
    it "works 4" do
      ticket = [38, 6, 12]
      rules = [
        Aoc2020::TicketRule.new(field: "class", ranges: [(1..3), (5..7)]),
        Aoc2020::TicketRule.new(field: "row", ranges: [(6..11), (33..44)]),
        Aoc2020::TicketRule.new(field: "seat", ranges: [(13..40), (45..50)]),
      ]

      Aoc2020.find_invalid_ticket_field(ticket, rules).should eq([12])
    end
  end
end
