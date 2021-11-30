class Integer
  def to_36_bit
    value = to_s(2)
    value.prepend('0' * (36 - value.size))
  end
end
