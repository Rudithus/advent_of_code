class Forward
  def initialize(value)
    @value = value
  end

  def move(waypoint, position)
    @value.times do
      position.x += waypoint.x
      position.y += waypoint.y
    end
    puts "Marching onwards! #{@value}, new position: #{position}"
  end
end

class North
  def initialize(value)
    @value = value
  end

  def move(waypoint, *)
    waypoint.y += @value
    puts "I'm gonna go North #{@value}, new waypoint: #{waypoint}"
  end
end

class South
  def initialize(value)
    @value = value
  end

  def move(waypoint, *)
    waypoint.y -= @value
    puts "Moving south by #{@value}, new waypoint: #{waypoint}"
  end
end

class East
  def initialize(value)
    @value = value
  end

  def move(waypoint, *)
    waypoint.x += @value
    puts "Time to move east #{@value}, new waypoint: #{waypoint}"
  end
end

class West
  def initialize(value)
    @value = value
  end

  def move(waypoint, *)
    waypoint.x -= @value
    puts "Lets go west for a while #{@value}, new waypoint #{waypoint}"
  end
end

class Right
  def initialize(value)
    @value = value
  end

  def move(waypoint, *)
    degree = @value % 360
    case degree
    when 90
      waypoint.x, waypoint.y = waypoint.y, -waypoint.x
    when 180
      waypoint.x = -waypoint.x
      waypoint.y = -waypoint.y
    when 270
      waypoint.y, waypoint.x = waypoint.x, -waypoint.y
    end
    puts "I wonder whats there on my right #{@value}, new waypoint #{waypoint}"
  end
end

class Left
  def initialize(value)
    @value = value
  end

  def move(waypoint, *)
    degree = @value % 360
    case degree
    when 270
      waypoint.x, waypoint.y = waypoint.y, -waypoint.x
    when 180
      waypoint.x = -waypoint.x
      waypoint.y = -waypoint.y
    when 90
      waypoint.y, waypoint.x = waypoint.x, -waypoint.y
    end
    puts "Suppose I should look at my left #{@value} new waypoint #{waypoint}"
  end
end
