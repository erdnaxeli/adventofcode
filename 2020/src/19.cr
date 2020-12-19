module Aoc2020
  struct FuturRegex
    def initialize(@pattern : String)
    end

    def to_s
      "#{@pattern}"
    end

    def to_s(io)
      io << @pattern
    end
  end

  struct Rules
    @rules = Hash(Int32, String | FuturRegex).new

    def initialize(input)
      input.lines.map(&.split(": ")).each do |x|
        number, rule = x
        @rules[number.to_i] = rule
      end
    end

    def [](index : Int32) : Regex
      pattern = get(index)
      Regex.new("^#{pattern}$")
    end

    def get(index : Int32) : FuturRegex
      rule = @rules[index]
      if rule.is_a?(FuturRegex)
        rule
      else
        pattern = rule.split('|').map do |group|
          group.split.reject(' ').compact_map do |c|
            if c == ""
              nil
            elsif c == %("a")
              "a"
            elsif c == %("b")
              "b"
            elsif c.to_i == index
              "(?&g#{index})"
            elsif @rules[c.to_i].is_a?(FuturRegex)
              "(?&g#{c.to_i})"
            else
              get(c.to_i)
            end
          end.join
        end.join('|')

        @rules[index] = FuturRegex.new("(?<g#{index}>#{pattern})")
      end
    end

    # Solution without regex, because my regex don't work for p2, I don't know why :'(
    #
    # Calling this method after `#[]` or `#get` will break.
    def match(input, rule_idx = 0, i = 0, skip_first_group = false)
      # pp input
      if input.size == 0
        0
      else
        rule = @rules[rule_idx].as(String)
        if rule == %("a")
          # puts "#{(['\t'] * i).join}rule #{rule_idx}, testing 'a'"
          if input[0] == 'a'
            1
          else
            0
          end
        elsif rule == %("b")
          # puts "#{(['\t'] * i).join}rule #{rule_idx}, testing 'b'"
          if input[0] == 'b'
            1
          else
            0
          end
        else
          matched_chars = 0
          groups = rule.split('|')
          if skip_first_group
            groups = groups[1..]
          end
          skip = 0

          while skip < 4
            groups.each do |group|
              # puts "#{(['\t'] * i).join}rule #{rule_idx}, testing group #{group}"
              matched_chars = 0

              group_rules = group.split(' ').reject("")
              group_rules.each_with_index(1) do |group_rule, group_index|
                # puts "subcall #{matched_chars}, #{input.size}, #{group_rule}, #{group}, #{group_index <= skip}"
                chars = match(
                  input[matched_chars..],
                  group_rule.to_i,
                  i + 1,
                  group_index <= skip,
                )
                if chars == 0
                  matched_chars = 0
                  break
                else
                  matched_chars += chars
                  if group_index < group_rules.size && matched_chars == input.size
                    # we matched too much
                    matched_chars = 0
                    break
                  end
                end
              end

              if matched_chars != 0
                break
              end
            end

            if matched_chars != 0
              break
            else
              skip += 1
            end
          end

          # puts "#{(['\t'] * i).join}rule #{rule_idx} matched #{matched_chars}"
          matched_chars
        end
      end
    end
  end

  def self.count_valid(messages, rule)
    messages.each_line.count do |l|
      if rule.matches?(l)
        puts l
        true
      else
        false
      end
    end
  end

  def self.count_valid_match(messages, rules)
    messages.each_line.count do |l|
      (rules.match(l).tap { |x| puts x } == l.size).tap { |x| x && puts l }
    end
  end

  INPUT_DAY19    = File.read("./inputs/19.txt")
  INPUT_DAY19_P2 = File.read("./inputs/19p2.txt")

  def self.day19p1
    rules, messages = INPUT_DAY19.split("\n\n")
    rules = Rules.new(rules)
    puts rules[0]

    count_valid(messages, rules[0])
  end

  def self.day19p2
    rules, messages = INPUT_DAY19_P2.split("\n\n")
    # rules = Rules.new(%(0: 1 | 1 0 1\n1: "a"))
    rules = Rules.new(rules)
    # puts rules[0]

    # count_valid(messages, rules[0])
    count_valid_match(messages, rules)
  end
end

# puts Aoc2020.day19p1
puts Aoc2020.day19p2
