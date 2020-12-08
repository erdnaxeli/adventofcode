require "digest/md5"

struct AdventCoin
  def initialize(@input : String)
  end

  def salt(x) : Int32
    zeros = "0"*x
    (0..).each do |i|
      h = Digest::MD5.hexdigest(@input + i.to_s)
      if h[0...x] == zeros
        return i
      end
    end
  end
end
