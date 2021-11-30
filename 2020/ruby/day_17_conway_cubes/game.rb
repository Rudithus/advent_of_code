class Game
  def initialize(rules, cubes)
    @rules = rules
    @cubes = cubes
  end

  def iterate
    rule_results = cubes.map { |cube| rules.map { |rule| rule.apply(cube) } }.flatten
    rule_results.each(&:execute)
  end
end
