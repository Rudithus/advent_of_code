load 'commands.rb'
load 'handheld.rb'

def input
  File.open('input.txt').readlines(&:chomp)
      .map { |line| line.split(' ') }
      .each_with_index.map { |(c, v), i| Object.const_get(c.capitalize).new(i, v.to_i) }
end

def change_command(command, commands)
  if command.instance_of?(Jmp)
    new_command = Nop.new(command.position, command.value)
  elsif command.instance_of?(Nop)
    new_command = Jmp.new(command.position, command.value)
  else
    raise StandardError
  end
  commands[command.position] = new_command
end

def execute?(handheld, commands)
  puts '--transaction start--'
  handheld.start_transaction
  handheld.execute(commands)
rescue HandheldError => e
  puts "bootup error on command #{e}"
  puts '--rollback--'
  handheld.rollback_transaction
else
  puts "bootup completed, acc: #{handheld.state.acc}"
  true
end

def debug(handheld, commands)
  last_command = handheld.undo_last
  until last_command.nil?
    unless last_command.instance_of?(Acc)
      change_command(last_command, commands)

      return true if execute?(handheld, commands)

      commands[last_command.position] = last_command
    end
    last_command = handheld.undo_last
  end
end

commands = input.to_a
handheld = Handheld.new

begin
  handheld.execute(commands)
rescue HandheldError => e
  puts "boot error on command: #{e}"

  puts '****** start debugging ******'
  debug(handheld, commands)
else
  puts "bootup completed, acc: #{handheld.state.acc}"
end
