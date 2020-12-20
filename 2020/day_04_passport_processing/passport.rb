load 'field.rb'

class Passport
  @@important_fields = %w[byr iyr eyr hgt hcl ecl pid]

  def initialize(line)
    @fields = line.split(/[\n, ]/).map { |l| Field.new(l) }
  end

  def valid?
    @@important_fields.all? { |f| @fields.map(&:key).include?(f) } && @fields.all?(&:valid?)
  end
end
