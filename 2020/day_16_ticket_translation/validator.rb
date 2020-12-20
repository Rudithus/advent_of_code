class Validator
  attr_reader :name

  def initialize(name, constraints)
    @name = name
    @constraints = constraints.map { |c| Constraint.new(*c) }
  end

  def invalid?(value)
    @constraints.all? { |c| !c.check(value) }
  end

  def to_s
    "validator: #{@name}, #{@constraints.join(',')}"
  end
end

class Constraint
  def initialize(min, max)
    @min = min
    @max = max
  end

  def check(value)
    value.between?(@min, @max)
  end

  def to_s
    "min: #{@min}, max: #{@max}"
  end
end
