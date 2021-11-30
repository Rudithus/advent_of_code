class CountryIdValidator
  def self.valid?(_value)
    true
  end
end

class BirthYearValidator
  def self.valid?(value)
    year = value.to_i
    year.between?(1920, 2002)
  end
end

class IssueYearValidator
  def self.valid?(value)
    year = value.to_i
    year.between?(2010, 2020)
  end
end

class ExpireYearValidator
  def self.valid?(value)
    year = value.to_i
    year.between?(2020, 2030)
  end
end

class HeightValidator
  def self.valid?(value)
    if value.include?('cm')
      length = value.split('cm')[0].to_i
      length.between?(150, 193)
    elsif value.include?('in')
      length = value.split('in')[0].to_i
      length.between?(59, 76)
    end
  end
end

class EyeColorValidator
  @valid_eyecolors = %w[amb blu brn gry grn hzl oth]
  def self.valid?(value)
    @valid_eyecolors.include?(value)
  end
end

class HairColorValidator
  def self.valid?(value)
    pattern = /^#[a-z,0-9]{6}$/
    pattern.match?(value)
  end
end

class PassportValidator
  def self.valid?(value)
    pattern = /^\d{9}$/
    pattern.match?(value)
  end
end
