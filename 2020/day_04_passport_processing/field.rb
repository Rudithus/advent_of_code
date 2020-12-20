load 'validators.rb'

class Field
  attr_reader :key, :value

  @@validators =
  {
    'byr' => BirthYearValidator,
    'iyr' => IssueYearValidator,
    'eyr' => ExpireYearValidator,
    'hgt' => HeightValidator,
    'hcl' => HairColorValidator,
    'ecl' => EyeColorValidator,
    'pid' => PassportValidator,
    'cid' => CountryIdValidator
  }

  def initialize(line)
    @key, @value = line.split(':')
  end

  def valid?
    @@validators[key].valid?(@value) if @@validators.key?(key)
  end
end
