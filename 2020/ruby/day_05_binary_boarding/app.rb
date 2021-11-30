def input
  file = File.open('input.txt').readlines.map(&:chomp)
end

def search(shift_rights)
  index = 0
  remaining = 2 ** shift_rights.count
  shift_rights.each do |shift|
    index += (remaining / 2) if shift
    remaining /= 2
  end
  index
end

def seat_number(line)
  column_bits = line[7..].each_char.map { |c| c == 'R' }
  row_bits = line[0..6].each_char.map { |c| c == 'B' }
  column = search(column_bits.to_a)
  row = search(row_bits.to_a)

  (row * 8) + column
end

seat_numbers = input.map { |l| seat_number(l)}.sort
min = seat_numbers.min
max = seat_numbers.max
arr = (min..max).to_a
puts arr - seat_numbers
