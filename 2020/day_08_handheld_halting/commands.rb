class Command
  attr_reader :position, :value

  def initialize(position, value)
    @position = position
    @value = value
  end

  def to_s
    "#{self.class.name} at line #{position}, value: #{value}"
  end
end

class Nop < Command
  def execute(state)
    state.pos += 1
  end

  def undo(state)
    state.pos -= 1
  end
end

class Jmp < Command
  def execute(state)
    state.pos += value
  end

  def undo(state)
    state.pos -= value
  end
end

class Acc < Command
  def execute(state)
    state.pos += 1
    state.acc += value
  end

  def undo(state)
    state.pos -= 1
    state.acc -= value
  end
end
