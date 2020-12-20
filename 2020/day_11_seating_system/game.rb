class Game
  def initialize(rules, seats)
    @seats = seats.flatten.reject(&:nil?).to_a
    @rules = rules
  end

  def occupied_seat_count
    @seats.count(&:occupied?)
  end

  def iterate
    results = @seats.map { |seat| @rules.map { |rule| rule.execute(seat) } }.flatten.reject(&:nil?)
    results.each(&:apply)
  end
end
