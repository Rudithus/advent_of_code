class RepeatingArray < Array
  def [](index)
    index = index % count
    super
  end
end
