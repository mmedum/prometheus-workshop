using Microsoft.AspNetCore.Mvc;
using Prometheus;

namespace dotnetservice.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class SampleController : ControllerBase
    {
        private static readonly Counter CallsCounter = 
        Metrics
        .CreateCounter(
        "dotnetservice_samplecontroller_counter", 
        "Number of calls to counter in sample controller.");

        private static readonly Gauge SampleGauge = Metrics
            .CreateGauge("dotnetservice_samplegauge", "Gauge set by api.");

        [HttpGet("")]
        public IActionResult Get(){
            return Ok();
        }

        [HttpGet("counter")]
        public IActionResult Counter(){
            CallsCounter.Inc();
            return Ok();
        }

        [HttpPost("gauge")]
        public IActionResult Gauge([FromForm] int absoluteNumber){
            SampleGauge.Set(absoluteNumber);
            return Ok();
        }
    }
}
