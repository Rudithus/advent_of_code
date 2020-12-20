def input
  file = File.open('input.txt')
  [file.readline.chomp.to_i, file.readline.split(',').map(&:to_i)]
end

arrival_time, time_table = input

@schedule = time_table.map.with_index.reject { |time, _| time.zero? }.sort_by { |time, _| time }

def check_number(time, shift, timestamp)
  if ((timestamp + shift) % time).zero?
    puts "found one #{timestamp}"
    true
  else
    false
  end
end

first = @schedule.first[0]
@schedule.drop(1).reduce([first, first]) do |(current, acc), (time, shift)|
  puts "new loop: current:#{current} acc:#{acc} time:#{time} shift:#{shift}"
  multiplier = 1
  multiplier += 1 until check_number(time, shift, current + (multiplier * acc))
  puts '-----'
  [current + multiplier * acc, acc * time]
end
