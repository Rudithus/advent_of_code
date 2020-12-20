class Policy
  def initialize(char, pos1, pos2)
    @char = char
    @pos1 = pos1 - 1
    @pos2 = pos2 - 1
  end

  def check(password)
    exists_pos1 = password[@pos1] == @char
    exists_pos2 = password[@pos2] == @char

    (exists_pos1 or exists_pos2) and !(exists_pos1 and exists_pos2)
  end
end
