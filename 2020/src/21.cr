module Aoc2020
  record Food, ingredients : Array(String), allergens = Array(String).new

  def self.read_input(input) : Array(Food)
    input.lines.map do |line|
      parts = line.split(" (contains ")
      ingredients = parts[0].split(' ')

      if parts.size == 2
        allergens = parts[1][...-1].split(", ")
        Food.new(ingredients: ingredients, allergens: allergens)
      else
        Food.new(ingredients: ingredients)
      end
    end
  end

  struct Solver
    @ingredients = Set(String).new

    getter allergens = Hash(String, Set(String)).new { |h, k| h[k] = Set(String).new }

    def initialize(@foods : Array(Food))
      find_potential_allergen
    end

    def find_potential_allergen
      @foods.each do |food|
        @ingredients.concat food.ingredients
        food.allergens.each do |allergen|
          if @allergens[allergen].size == 0
            @allergens[allergen] = Set(String).new(food.ingredients)
          else
            @allergens[allergen] &= Set(String).new(food.ingredients)
          end
        end
      end
    end

    def non_allergen
      ingredients_with_allergens = @allergens.values.reduce { |acc, x| acc += x }
      @ingredients - ingredients_with_allergens
    end

    def count_apparitions(ingredients)
      ingredients.map do |ingredient|
        @foods.count do |food|
          food.ingredients.includes?(ingredient)
        end
      end.sum
    end

    def find_valid_allergen
      done = Set(String).new

      while done.size != @allergens.size
        allergen = @allergens.keys.find do |k|
          !done.includes?(k) && @allergens[k].size == 1
        end.not_nil!
        ingredient = @allergens[allergen].first
        @allergens.reject { |k, _| k == allergen }.each do |_, ingredients|
          ingredients.delete(ingredient)
        end

        done << allergen
      end
    end
  end

  INPUT_DAY21 = File.read("./inputs/21.txt")

  def self.day21p1
    foods = read_input(INPUT_DAY21)
    solver = Solver.new(foods)
    solver.find_potential_allergen
    solver.count_apparitions(solver.non_allergen)
  end

  def self.day21p2
    foods = read_input(INPUT_DAY21)
    solver = Solver.new(foods)
    solver.find_potential_allergen
    solver.find_valid_allergen
    solver.allergens.to_a.sort_by { |k, v| k }.map { |k, v| v.first }.join(',')
  end
end

# puts Aoc2020.day21p2
