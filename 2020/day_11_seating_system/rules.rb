class RuleResult
  def initialize(result)
    @result = result
  end

  def apply
    @result.call
  end
end

class OccupyIfAloneRule
  def execute(seat)
    RuleResult.new(-> { seat.occupy }) if seat.neighbours.count(&:occupied?).zero?
  end
end

class VacateIfCrowdedRule
  def execute(seat)
    RuleResult.new(-> { seat.vacate }) if seat.neighbours.count(&:occupied?) >= 5
  end
end
