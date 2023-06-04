require 'gosu'

class SimulationWindow < Gosu::Window
  def initialize
    super(800, 600)
    self.caption = "地震波シミュレーション"
    
    # シミュレーションのパラメータ
    @L = 100.0   # シミュレーション領域の長さ（km）
    @dx = 1.0    # メッシュのステップサイズ（km）
    @dt = 0.001  # 時間のステップサイズ（秒）
    @nt = 1000   # シミュレーションのステップ数
    
    # メッシュの設定
    @nx = (@L / @dx).to_i + 1
    @x = Array.new(@nx) { |i| i * @dx }
    
    # 地盤の初期条件
    @vs = 2000.0 # 地盤のS波速度（m/s）
    @rho = 2000.0 # 地盤の密度（kg/m^3）
    @h = Array.new(@nx, 1000.0)
    
    # 応力・速度・変位の初期条件
    @stress = Array.new(@nx, 0.0)
    @velocity = Array.new(@nx, 0.0)
    @displacement = Array.new(@nx, 0.0)
    
    # シミュレーションの実行
    @nt.times do |n|
      # 応力の更新
      @stress[1..-2] += (@vs**2 * @dt / @dx) * (@velocity[2..-1] - @velocity[0..-3])
      
      # 速度の更新
      @velocity[1..-2] += (@dt / @rho / @h[1..-2]) * (@stress[2..-1] - @stress[0..-3])
      
      # 変位の更新
      @displacement[1..-2] += @dt * @velocity[1..-2]
      
      # 境界条件（自由端境界）
      @velocity[0] = 0.0
      @velocity[-1] = 0.0
    end
  end
  
  def draw
    scale = self.width / @L
    
    # 変位を描画
    @nx.times do |i|
      x = @x[i] * scale
      y = self.height / 2 + @displacement[i] * scale
      self.draw_line(x, self.height / 2, Gosu::Color::RED, x, y, Gosu::Color::RED)
    end
  end
end

window = SimulationWindow.new
window.show
