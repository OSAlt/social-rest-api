package social.service.impl;

import lombok.NoArgsConstructor;
import org.apache.commons.io.IOUtils;
import org.apache.http.HttpEntity;
import org.apache.http.StatusLine;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.powermock.core.classloader.annotations.PrepareForTest;
import org.powermock.modules.junit4.PowerMockRunner;
import social.service.SocialDataService;

import java.io.IOException;

import static org.junit.Assert.assertEquals;

@RunWith(PowerMockRunner.class)
@PrepareForTest({HttpClients.class, EntityUtils.class})
@NoArgsConstructor
public class InstagramSocialDataServiceImplTest {

    SocialDataService service;

    @Before
    public void setup() throws IOException {
        service= new InstagramSocialDataServiceImpl("dummy");
        SocialDataHelperTest.setupMockedObjects("/responses/instagram_count.json");

    }

    @Test
    public void getCount() {
        int expected = 4807;
        int count = service.getCount();
        assertEquals(expected,count);
    }

    @Test(expected = UnsupportedOperationException.class)
    public void authorize() {
        service.authorize();
    }

    @Test(expected = UnsupportedOperationException.class)
    public void saveRequestToken() {
        service.saveRequestToken("dummy", "value");
    }

}