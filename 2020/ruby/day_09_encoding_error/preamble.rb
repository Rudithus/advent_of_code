class Preamble
  def initialize(numbers)
    @numbers = numbers
    @size = numbers.count

    sums = numbers.combination(2).map(&:sum)
    @sum_hash = counted_dictionary(sums)
    @sum_table = stacked_table(sums, @size)
  end

  private

  def counted_dictionary(array)
    array.each_with_object(Hash.new(0)) { |sum, hash| hash[sum] += 1 }
  end

  def stacked_table(array, stack_size)
    (1..stack_size - 1).each_with_object([]) { |i, table| table[i - 1] = array.shift(stack_size - i) }
  end

  def remove_sum(sum)
    @sum_hash[sum] -= 1
    @sum_hash.delete(sum) if @sum_hash[sum].zero?
  end

  def remove_first_number
    @sum_table[0].each { |s| remove_sum(s) }
    @sum_table = @sum_table.drop(1)
    @numbers = @numbers.drop(1)
  end

  def add_new_sum(sum, row)
    @sum_hash[sum] += 1
    if @sum_table[row].nil?
      @sum_table[row] = [sum]
    else
      @sum_table[row].push(sum)
    end
  end

  public

  def add(number)
    remove_first_number

    @numbers.map { |n| n + number }.each_with_index do |sum, row|
      add_new_sum(sum, row)
    end

    @numbers.push(number)
  end

  def include?(number)
    @sum_hash.key?(number)
  end

  def to_s
    "#{@numbers} - #{@sum_table} - #{@set}"
  end
end
