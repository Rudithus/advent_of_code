class Bag
  attr_reader :parents, :name, :children

  def initialize(name)
    @name = name
    @parents = []
    @children = []
  end

  def add_parent(bag)
    @parents.push(bag)
  end

  def add_child(bag)
    @children.push(bag)
  end
end
