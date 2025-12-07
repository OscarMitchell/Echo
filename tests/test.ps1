$client = New-Object System.Net.Sockets.TcpClient("localhost", 23)
$stream = $client.GetStream()
$writer = New-Object System.IO.StreamWriter($stream)
$writer.writeLine("Closing connection in...")
$writer.flush()
$writer.writeLine("5...")
$writer.flush()
Start-Sleep -Seconds 1
$writer.writeLine("4...")
$writer.flush()
Start-Sleep -Seconds 1
$writer.writeLine("3...")
$writer.flush()
Start-Sleep -Seconds 1
$writer.writeLine("2...")
$writer.flush()
Start-Sleep -Seconds 1
$writer.writeLine("1...")
$writer.flush()
Start-Sleep -Seconds 1
$client.Close()
