load 'validator.rb'
load 'ticket.rb'

def input
  file = File.open('input.txt')
  validators = []
  my_ticket = nil
  nearby_tickets = []

  while (line = read_until_empty_line(file))
    name, constraints = line.split(':')
    validator = Validator.new(name, split_constraint_line(constraints))
    validators.push(validator)
  end

  file.readline
  while (line = read_until_empty_line(file))
    my_ticket = create_ticket(line, validators)
  end

  file.readline
  while (line = read_until_empty_line(file))
    nearby_tickets.push(create_ticket(line, validators))
  end

  [nearby_tickets, validators, my_ticket]
end

def read_until_empty_line(file)
  line = file.gets
  line.nil? || line == "\n" ? nil : line.chomp
end

def split_constraint_line(line)
  line.split(' or ').map { |c| c.split('-').map(&:to_i) }
end

def create_ticket(line, validators)
  values = line.split(',').map(&:to_i)
  Ticket.new(values, validators)
end

def ticket_scanning_error_rate(tickets)
  tickets.map(&:invalid_numbers).flatten.reject(&:nil?).sum
end

@tickets, @validators, @my_ticket = input

puts ticket_scanning_error_rate(@tickets)

valid_values = @tickets.reject(&:invalid?).map(&:values).transpose
@correct_fields = Hash.new { |h, k| h[k] = [] }

@validators.each do |validator|
  valid_values.each_with_index do |value_array_at_index, index|
    next unless value_array_at_index.all? { |value| !validator.invalid?(value) }

    arr = @correct_fields[index]
    arr.push(validator.name)
    @correct_fields[index] = arr
  end
end

@correct_fields.sort_by { |_k, v| v.count }.reverse
               .each_cons(2) { |(k1, v1), (_, v2)| @correct_fields[k1] = (v1 - v2) }

puts @correct_fields.select { |_k, v| v[0].start_with?('departure') }.map { |k, _v| @my_ticket.values[k] }.reduce(&:*)
