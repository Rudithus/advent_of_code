class Node
  attr_accessor :value

  def initialize(value)
    @value = value
    @children = []
    @branch_count = nil
  end

  def add_connection(node)
    @children.push(node)
  end

  def branch_count
    return @branch_count unless @branch_count.nil?
    return 1 if @children.count.zero?

    @branch_count = @children.sum(&:branch_count)
    @branch_count
  end
end
