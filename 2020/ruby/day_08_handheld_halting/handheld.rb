require 'set'

class Handheld
  attr_reader :state

  def initialize
    @state = State.new
    @command_history = []
    @command_set = Set.new
    @in_transaction = false
  end

  private

  def execute_command(command)
    puts command
    command.execute(state)
    @transaction_commands += 1 if @in_transaction
    @command_history.push(command)
  end

  def undo_command(command)
    puts "undo #{command}"
    command.undo(state)

    @command_set.delete(command)

    command
  end

  public

  def execute(instructions)
    command = instructions[state.pos]
    until command.nil?
      raise HandheldError, command unless @command_set.add?(command)

      execute_command(command)

      command = instructions[state.pos]
    end
  end

  def undo_last
    command = @command_history.pop
    return if command.nil?

    undo_command(command)
  end

  def start_transaction
    @in_transaction = true
    @transaction_commands = 0
  end

  def end_transaction
    @transaction_commands = 0
    @in_transaction = false
  end

  def rollback_transaction
    @transaction_commands.times do
      undo_last
    end
    end_transaction
  end
end

class State
  attr_accessor :command_history, :pos, :acc

  def initialize
    @pos = 0
    @acc = 0
  end
end

class HandheldError < StandardError
  attr_accessor :command

  def initialize(command)
    @command = command
    super
  end
end
