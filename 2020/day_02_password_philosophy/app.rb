load 'policy.rb'

def get_passwords
  File.open('input.txt').readlines.map(&:chomp)
end

def parse_policy(line)
  (min_max, char) = line.split(' ')
  (min, max) = min_max.split('-')
  Policy.new(char, min.to_i, max.to_i)
end

def seperate_policy_from_password(line)
  (policy, password) = line.split(':').map(&:strip)
end

passwords = get_passwords
puts passwords
  .map { |line| seperate_policy_from_password(line) }
  .map { |policy, password| parse_policy(policy).check(password) }
  .count(true)