# seismic-wave-simulation

2次元差分法の地震波シミュレーション

## シミュレーションの実行
```
go run main.go
```
## シミュレーションのビジュアル化

### 実行準備
```
pip3 install matplotlib pandas
```

### 実行
```
python main.py
```

## mp4動画の作成
```
cd images; ffmpeg -framerate 30 -i frame_%04d.png -c:v libx264 -pix_fmt yuv420p output.mp4
```

### 全ての処理を実行
```
./command.sh
```
