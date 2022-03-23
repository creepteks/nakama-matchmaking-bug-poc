using Nakama;
using System.Collections.Generic;
using UnityEngine;

public class GameController : MonoBehaviour
{
    public ConnectionVars connection;
    private ISocket socket;

    // Start is called before the first frame update
    private async void Start()
    {
        // Use "https" scheme if you've setup SSL.
        var client = new Client("http", connection.ip, connection.port, "defaultkey");

        // Should use a platform API to obtain a device identifier.
#if UNITY_EDITOR
        var deviceId = "EDITOR-instance";

#elif UNITY_STANDALONE

        var deviceId = System.Guid.NewGuid().ToString();
#endif
        var session = await client.AuthenticateDeviceAsync(deviceId);
        System.Console.WriteLine("New user: {0}, {1}", session.Created, session);

        socket = client.NewSocket();
        await socket.ConnectAsync(session);

        socket.ReceivedMatchmakerMatched += Socket_ReceivedMatchmakerMatched;

        var query = "*";
        var minCount = 2;
        var maxCount = 4;

        var stringProperties = new Dictionary<string, string>() {
            {"region", "europe"}
        };

        var matchmakerTicket = await socket.AddMatchmakerAsync(query, minCount, maxCount, stringProperties);
    }

    private void Socket_ReceivedMatchmakerMatched(IMatchmakerMatched obj)
    {
        socket.JoinMatchAsync(obj.MatchId);
    }

    private void OnDestroy()
    {
        socket.CloseAsync();
    }
}
