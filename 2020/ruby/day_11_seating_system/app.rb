load 'seat.rb'
load 'game.rb'
load 'rules.rb'

def input
  File.readlines('input.txt').map(&:chomp).map(&:chars)
end

def top_neighbour(row, column)
  until row.zero?
    row -= 1
    return @seats[row][column] unless @seats[row][column].nil?
  end
end

def top_left_neighbour(row, column)
  until row.zero? || column.zero?
    row -= 1
    column -= 1
    return @seats[row][column] unless @seats[row][column].nil?
  end
end

def left_neighbour(row, column)
  until column.zero?
    column -= 1
    return @seats[row][column] unless @seats[row][column].nil?
  end
end

def top_right_neighbour(row, column)
  until row.zero?
    row -= 1
    column += 1
    return @seats[row][column] unless @seats[row][column].nil?
  end
end

def neighbours(row, column)
  top = top_neighbour(row, column)
  left = left_neighbour(row, column)
  top_left = top_left_neighbour(row, column)
  top_right = top_right_neighbour(row, column)
  [top, left, top_left, top_right]
end

def print_room
  @seats.each do |row|
    a = row.map { |s| s.nil? ? '.' : s }
    puts a.join('')
  end
end

@rules = [OccupyIfAloneRule.new, VacateIfCrowdedRule.new]
@seats = []

input.each_with_index do |line, row|
  @seats[row] = []
  line.each.with_index do |char, column|
    if char == '.'
      @seats[row][column] = nil
      next
    end

    seat = Seat.new(Position.new(row, column))
    @seats[row][column] = seat

    neighbours(row, column).reject(&:nil?).each { |neighbour| seat.add_neighbour(neighbour) }
  end
end

@game = Game.new(@rules, @seats)

current_count = -1

until @game.occupied_seat_count == current_count
  current_count = @game.occupied_seat_count
  @game.iterate
end

print_room

puts @game.occupied_seat_count
