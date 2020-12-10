require "file"

module Aoc2020
  struct Passport
    property byr : String? = nil
    property iyr : String? = nil
    property eyr : String? = nil
    property hgt : String? = nil
    property hcl : String? = nil
    property ecl : String? = nil
    property pid : String? = nil
    property cid : String? = nil

    def has_fields?
      !byr.nil? && !iyr.nil? && !eyr.nil? && !hgt.nil? && !hcl.nil? && !ecl.nil? && !pid.nil?
    end

    def is_valid?
      has_fields? && byr_valid? && iyr_valid? && eyr_valid? && hgt_valid? && hcl_valid? && ecl_valid? && pid_valid?
    end

    def byr_valid?
      number_valid?(1920, 2002, byr)
    end

    def iyr_valid?
      number_valid?(2010, 2020, iyr)
    end

    def eyr_valid?
      number_valid?(2020, 2030, eyr)
    end

    def hgt_valid?
      if result = /(?<height>\d+)(?<unit>cm|in)/.match(hgt.not_nil!)
        if result["unit"] == "cm"
          number_valid?(150, 193, result["height"])
        else
          number_valid?(59, 76, result["height"])
        end
      else
        false
      end
    end

    def hcl_valid?
      /^#[0-9a-f]{6}$/.matches? hcl.not_nil!
    end

    def ecl_valid?
      %w(amb blu brn gry grn hzl oth).includes?(ecl)
    end

    def pid_valid?
      /^\d{9}$/.matches? pid.not_nil!
    end

    private def number_valid?(to, from, value)
      (to..from).includes?(value.not_nil!.to_i { 0 })
    end
  end

  def self.read_passports(input)
    passports = Array(Passport).new

    input.split(/\n\n/m) do |passport_input|
      passport = Passport.new

      passport_input.split(/\s/m, remove_empty: true) do |field_input|
        field, value = field_input.split(':')

        {% begin %}
          case field
            {% for var in Passport.instance_vars %}
              when {{var.stringify}}
                passport.{{var.id}} = value
            {% end %}
          end
        {% end %}
      end

      passports << passport
    end

    passports
  end

  def self.day4p1
    puts read_passports(File.read("./inputs/04.txt")).count &.has_fields?
  end

  def self.day4p2
    puts read_passports(File.read("./inputs/04.txt")).count &.is_valid?
  end
end
