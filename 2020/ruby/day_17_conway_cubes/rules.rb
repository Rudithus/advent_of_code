class DeactivateIfLonely
  def apply(cube)
    RuleResult.new(-> { cube.deactivate }) if cube.active_neighbour_count < 2
  end
end

class DeactivateIfCrowded
  def apply(cube)
    RuleResult.new(-> { cube.deactivate }) if cube.active_neighbour_count > 3
  end
end

class ActivateIfThreeNeighbours
  def apply(cube)
    RuleResult.new(-> { cube.activate }) if cube.active_neighbour_count == 3
  end
end

class RuleResult
  def initialize(result)
    @result = result
  end

  def execute
    @result.call
  end
end
