class Node
  attr_reader :value

  def initialize(value)
    @value = value
    @children = []
    @branch_count = nil
  end

  def add_connection(node)
    @children.push(node)
  end

  def can_connect?(node)
    node.value - @value <= 3
  end

  def branch_count
    return @branch_count unless @branch_count.nil?

    @branch_count = @children.count.zero? ? 1 : @children.sum(&:branch_count)
  end
end
