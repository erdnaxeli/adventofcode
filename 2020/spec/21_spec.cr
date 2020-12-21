require "./spec_helper"
require "../src/21"

input = %(mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
)

describe Aoc2020 do
  describe ".find_non_allergen_ingredients" do
    it "works" do
      foods = Aoc2020.read_input(input)
      non_allergens = Aoc2020.find_non_allergen_ingredients(foods)
      non_allergens.should eq(Set{"kfcds", "nhms", "sbzzf", "trh"})
    end
  end

  describe ".count_apparitions" do
    it "works" do
      foods = Aoc2020.read_input(input)
      apparitions = Aoc2020.count_apparitions(foods, Set{"kfcds", "nhms", "sbzzf", "trh"})
      apparitions.should eq(5)
    end
  end
end
