class Ticket
  attr_reader :values

  def initialize(values, validators)
    @values = values
    @validators = validators
  end

  def invalid?
    @values.any? { |value| @validators.all? { |validator| validator.invalid?(value) } }
  end

  def invalid_numbers
    @values.select { |value| @validators.all? { |validator| validator.invalid?(value) } }
  end

  def to_s
    @values.join(',')
  end
end
